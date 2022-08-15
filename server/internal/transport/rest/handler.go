package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"Skavengerr/coins-crud/internal/types"
	"Skavengerr/coins-crud/util"

	"github.com/gorilla/mux"
)

type Coins interface {
	CreateCoin(coin types.Coin) error
	GetCoinByID(id int64) (types.Coin, error)
	GetAllCoins() ([]types.Coin, error)
	DeleteCoin(id int64) error
}

type Handler struct {
	coinsService Coins
}

func NewHandler(coins Coins) *Handler {
	return &Handler{
		coinsService: coins,
	}
}

func (h *Handler) InitRouter() *mux.Router {
	r := mux.NewRouter()

	coins := r.PathPrefix("/api/coins").Subrouter()
	{
		coins.HandleFunc("/create", h.createCoin).Methods(http.MethodPost)
		coins.HandleFunc("", h.getAllCoins).Methods(http.MethodGet)
		coins.HandleFunc("/{id:[0-9]+}", h.getCoinByID).Methods(http.MethodGet)
		coins.HandleFunc("/{id:[0-9]+}", h.deleteCoin).Methods(http.MethodDelete)
	}

	return r
}

func (h *Handler) getAllCoins(w http.ResponseWriter, r *http.Request) {
	coins, err := h.coinsService.GetAllCoins()
	if err != nil {
		log.Println("getAllCoins() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(coins)
	if err != nil {
		log.Println("getAllCoins() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}

func (h *Handler) getCoinByID(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		log.Println("getCoinByID() error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	coin, err := h.coinsService.GetCoinByID(id)
	if err != nil {
		util.CheckErr(err)

		log.Println("getCoinByID() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(coin)
	if err != nil {
		log.Println("getCoinByID() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}

func (h *Handler) createCoin(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var coin types.Coin
	if err = json.Unmarshal(reqBytes, &coin); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.coinsService.CreateCoin(coin)
	if err != nil {
		log.Println("createCoin() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) deleteCoin(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		log.Println("deleteCoin() error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.coinsService.DeleteCoin(id)
	if err != nil {
		log.Println("deleteCoin() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getIdFromRequest(r *http.Request) (int64, error) {
	fmt.Println("r", r)
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return 0, err
	}

	if id == 0 {
		return 0, errors.New("id can't be 0")
	}

	return id, nil
}
