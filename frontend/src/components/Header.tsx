import React from "react";
import { ethers } from "ethers";
import useIsMetaMaskInstalled from "../useIsMetaMaskInstalled";
import FavoriteIcon from '@mui/icons-material/Favorite';
import Grid from '@mui/material/Grid';


interface Props {
  setAccount: React.Dispatch<React.SetStateAction<string | undefined>>;
  account?: string;
  contractAddress: string;
  appName: string;
  network: string;
}

const Headers = ({ setAccount, account, contractAddress, appName, network}: Props) => {
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
    <header className="header">
      <Grid container>
        <Grid item xs={6} md={6}>
            <div className="title">
                <h1><FavoriteIcon/> {appName}</h1>
            </div>
        </Grid>
        <Grid item xs={6} md={6} className="wallet">
            {account && (
            <>
                <b>Connected as:</b>
                <small><a target="blank" href={"http://"+ network +".etherscan.io/address/" + account}> {account}</a></small>
            </>
            )}
            {!account && (
            <button onClick={handleOnConnect} disabled={!isMetaMaskInstalled}>
                Connect your wallet
            </button>
            )}
            {!isMetaMaskInstalled && <p>Please install MetaMask</p>}
        </Grid>
      </Grid>
    </header>
  );
};

export default Headers;
