import { ethers } from 'ethers';
import { WeddingDetails } from "../types";
import React, { useEffect, useState } from 'react';
import CheckCircleIcon from '@mui/icons-material/CheckCircle';
import { Avatar } from '@mui/material';
import { Grid } from '@mui/material';
import { Button } from '@mui/material';
import Attestation from './Attestation';


interface Props {
  weddingContract: ethers.Contract | undefined;
  account?: string;
}


const Ceremony = ({ weddingContract, account }: Props) => {
  const [wallet, setWallet] = useState<ethers.providers.JsonRpcSigner | null>(null);
  const [fiance1Status, setFiance1Status] = useState(false);
  const [fiance2Status, setFiance2Status] = useState(false);

  useEffect(() => {
    if (window.ethereum) {
      const provider = new ethers.providers.Web3Provider(window.ethereum);
      const signer = provider.getSigner();
      setWallet(signer);

      if (weddingContract) {
        (async () => {
          try {
            const [fiance1WalletAddr,] = await weddingContract.getFiance1();
            const [fiance2WalletAddr,] = await weddingContract.getFiance2();
            weddingContract.on("NewSignature", (from: string, timestamp: ethers.BigNumber, event: ethers.Event) => {
              if (from.toLowerCase() === fiance1WalletAddr.toLowerCase()) {
                setFiance1Status(true);
              } else if (from.toLowerCase() === fiance2WalletAddr.toLowerCase()) {
                setFiance2Status(true);
              }
            });
          } catch (err) {
            console.error(err);
          }
        })();
      }
    } else {
      console.log('Please install MetaMask!');
    }

    return () => {
      if (weddingContract) {
        weddingContract.removeAllListeners("NewSignature");
      }
    };
  }, [weddingContract]);

  const sayYes = async () => {
    if (wallet && weddingContract) {
      const contractWithSigner = weddingContract.connect(wallet);
      const tx = await contractWithSigner.sayYes();
      await tx.wait();
    }
  };

  return (
    <div>
      <Grid container id="greetings">
          <h1>Welcome everyone!</h1>
      </Grid>
      <Grid container>
        <Grid item xs={6} md={6} >
        

          <div className="weddingCard">
            <h2>Partner 1</h2>
            <p>Some description about me</p>
            <Avatar
              className='weddingAvatar'
              alt="fiance"
              src="https://bafybeiebtleqr4mtvuq273umq3v637mksfloul2zsfdlvxuwp6hmvubrie.ipfs.dweb.link/"
              sx={{ width: 110, height: 110 }}
            />
            <hr></hr>
            <Button className='weddingButton' disabled={fiance1Status} onClick={sayYes} variant="contained"  color="success" endIcon={<CheckCircleIcon />}>
              Say Yes
            </Button>
          </div>


        </Grid>
        <Grid item xs={6} md={6}>

        <div className="weddingCard">
          <h2>Partner 1</h2>
          <p>Some description about me</p>
          <Avatar
            className='weddingAvatar'
            alt="fiance"
            src="https://bafybeib4ndy3fuysk7fcjo6yxycdans4zbyzpnmajwfo45n7ufyd2aywj4.ipfs.dweb.link/"
            sx={{ width: 110, height: 110 }}
          />
          <hr></hr>
          <Button className='weddingButton' disabled={fiance2Status} onClick={sayYes} variant="contained"  color="success" endIcon={<CheckCircleIcon />}>
            Say Yes
          </Button>
        </div>

        <Attestation account={account} weddingContract={weddingContract} />
        </Grid>
      </Grid>
      
      {/* Today, {name1} and {name2} are getting married! */}
    </div>
  );
};

export default Ceremony;
