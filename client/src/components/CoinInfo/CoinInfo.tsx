import React, { useEffect, useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import axios from 'axios';
import { Button } from 'antd';

import { Coin } from 'shared/interfaces/coins';

export const CoinInfo: React.FC = () => {
  const { id } = useParams();
  const navigate = useNavigate();
  const [coinData, setCoinData] = useState<Coin>();

  const backToList = () => {
    navigate('/')
  }

  useEffect(() => {
    axios.get(`/api/coins/${id}`).then((res) => setCoinData(res.data));
  }, []);

  return (
    <div>
      <Button onClick={backToList}>Back</Button>
      <h2>Id: {coinData?.id}</h2>
      <h2>Name: {coinData?.name}</h2>
      <h2>Amount: {coinData?.amount}</h2>
    </div>
  );
};
