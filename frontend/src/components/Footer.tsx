import React from "react";
import { ethers } from "ethers";
import useIsMetaMaskInstalled from "../useIsMetaMaskInstalled";

interface Props {
  setAccount: React.Dispatch<React.SetStateAction<string | undefined>>;
  account?: string;
  contractAddress: string;
  appName: string;
}

const Headers = ({ setAccount, account, contractAddress, appName }: Props) => {
  const isMetaMaskInstalled = useIsMetaMaskInstalled();

  const handleOnConnect = () => {
    window.ethereum
      .request({ method: "eth_requestAccounts" })
      .then((accounts: string[]) => {
        setAccount(ethers.utils.getAddress(accounts[0]));
      })
      .catch((err: any) => console.log(err));
  };

  return (
    <footer className="footer">
      <p>
        Contract Address: <a target="blank" href="http://etherscan.io/address/{contractAddress}"> {contractAddress}</a>
      </p>
    </footer>
  );
};

export default Headers;
