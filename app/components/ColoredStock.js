import React from 'react';

export default function ColoredStock (props){

  let whichColor

  if ((props.openingPrice) < props.cell.price){
  	whichColor = "green"
  }
  else if(props.openingPrice > props.cell.price){
  	whichColor = "red"
  }
  else{
  	whichColor = "grey"
  }

  return (

    <div style={{color:whichColor}}>
    	{props.cell.symbol}
    </div>
  );
};