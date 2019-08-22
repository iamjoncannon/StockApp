import React from 'react';

export default function TransHistoryItem (props){

  const { data } = props;

  return (
    <div>
    	<span>{data.Symbol} </span>
    	<span>{data.Quantity} </span>
    	<span>{data.Date} </span>
    </div>
  );
};