import React from "react";
import { ethers } from "ethers";

interface Props {
  contractAddress: string;
}

const Headers = ({ contractAddress }: Props) => {

  return (
    <footer className="footer">
      <p>
        Contract Address: <a target="blank" href={"http://sepolia.etherscan.io/address/" + contractAddress}> {contractAddress}</a>
      </p>
    </footer>
  );
};

export default Headers;
