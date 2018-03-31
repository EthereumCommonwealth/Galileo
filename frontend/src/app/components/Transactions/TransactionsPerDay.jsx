import React from 'react';
import { Line } from 'react-chartjs-2';

const data = {
  labels: ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday'],
  datasets: [
    {
      label: 'Transactions per day',
      lineTension: 0.1,
      backgroundColor: 'rgba(75,192,192,0.4)',
      borderColor: 'rgba(75,192,192,1)',
      borderCapStyle: 'butt',
      borderDash: [],
      borderDashOffset: 0.0,
      borderJoinStyle: 'miter',
      pointBorderColor: 'rgba(75,192,192,1)',
      pointBackgroundColor: '#fff',
      pointBorderWidth: 1,
      pointHoverRadius: 5,
      pointHoverBackgroundColor: 'rgba(75,192,192,1)',
      pointHoverBorderColor: 'rgba(220,220,220,1)',
      pointHoverBorderWidth: 2,
      pointRadius: 1,
      pointHitRadius: 10,
      data: [990, 5000, 1000, 15000, 22400, 38900, 24000, 44000]
    }
  ]
};

const TransactionsPerDay = () => (
  <div className='TransactionsPerDay'>
    <h2 className='TransactionsPerDay-title'>Transactions per day</h2>
    <div className='TransactionsPerDay-content'>
      <Line data={data} />
    </div>
  </div>
)

export default TransactionsPerDay;
