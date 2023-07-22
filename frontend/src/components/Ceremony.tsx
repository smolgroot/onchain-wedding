import React, { useEffect, useState } from "react";
import { ethers } from "ethers";
// import { Signer } from "../types";


interface Props {
  account?: string;
  chatContract: ethers.Contract | undefined;
}


const Ceremony = ({ account, chatContract }: Props) => {
  return (
    <div className="photos">
      Ceremony
    </div>
  );
};

export default Ceremony;
