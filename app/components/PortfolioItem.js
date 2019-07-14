import React from 'react';

export default function PortfolioItem (props){

	const { data } = props

  return (
    <div>
    	<span>{data.symbol} </span>
    	<span>{data.quantity} </span>
    	<span>{data.price} </span>
    </div>
  );
};