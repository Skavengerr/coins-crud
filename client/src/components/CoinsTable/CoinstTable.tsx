import { Table } from 'antd';
import type { ColumnsType } from 'antd/es/table';
import axios from 'axios';
import { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import { Coin } from 'shared/interfaces/coins';

export function CoinsTable() {
  const [coinsData, setCoinsData] = useState([]);

  const columns: ColumnsType<Coin> = [
    {
      title: 'ID',
      dataIndex: 'id',
      key: 'id',
      render: (id) => <p>{id}</p>,
    },
    {
      title: 'Name',
      dataIndex: 'name',
      key: 'name',
      render: (name) => <Link to={`/${getIdByName(name)}`}>{name}</Link>,
    },
    {
      title: 'Amount',
      dataIndex: 'amount',
      key: 'amount',
      render: (amount) => <p>{amount}</p>,
    },
  ];

  const getIdByName = (name: string) => {
    // @ts-ignore
    return coinsData.find((coin: Coin) => coin.name === name)?.id;
  };

  useEffect(() => {
    axios.get('/api/coins').then((res) => setCoinsData(res.data));
  }, []);

  return (
    <div>
      <Table columns={columns} dataSource={coinsData} />
    </div>
  );
}
