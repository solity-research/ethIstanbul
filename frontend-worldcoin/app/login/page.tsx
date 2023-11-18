"use client"
import { useEffect, useState } from "react";
import { CredentialType, IDKitWidget, ISuccessResult,useIDKit } from '@worldcoin/idkit'
import { useRouter } from "next/navigation";
import { ethers } from "ethers";
import detectEthereumProvider from "@metamask/detect-provider";

declare global {
    interface Window {
        ethereum?: any;
    }
}

export const formatBalance = (rawBalance: string) => {
    const balance = (parseInt(rawBalance) / 1000000000000000000).toFixed(2)
    return balance
}

export const formatChainAsNum = (chainIdHex: string) => {
    const chainIdNum = parseInt(chainIdHex)
    return chainIdNum
}

const Login = () => {
    const [wordCoinAddress, setWordCoinAddress] = useState<ISuccessResult>()
    const [userWalletAddress, setUsetWalletAddress] = useState("")
    const { open, setOpen } = useIDKit()
    const [hasProvider, setHasProvider] = useState<boolean | null>(null)
    const initialState = { accounts: [], balance: "", chainId: "" }
    const [wallet, setWallet] = useState(initialState)
    const [isConnecting, setIsConnecting] = useState(false)
    const [error, setError] = useState(false)
    const [errorMessage, setErrorMessage] = useState("")
    
    useEffect(()=>{
        const refreshAccounts = (accounts: any) => {
            if (accounts.length > 0) {
                setUsetWalletAddress(accounts)
            } else {
                // if length 0, user is disconnected
                setWallet(initialState)
            }
        }

        const refreshChain = (chainId: any) => {
            setWallet((wallet) => ({ ...wallet, chainId }))
        }

        const getProvider = async () => {
            const provider = await detectEthereumProvider({ silent: true })
            setHasProvider(Boolean(provider))

            if (provider) {
                const accounts = await window.ethereum.request(
                    { method: 'eth_accounts' }
                )
                refreshAccounts(accounts)
                window.ethereum.on('accountsChanged', refreshAccounts)
                window.ethereum.on("chainChanged", refreshChain)
            }
        }
        const handleConnect = async () => {
            setIsConnecting(true)
            await window.ethereum.request({
                method: "eth_requestAccounts",
            })
                .then((accounts:Array<string>) => {
                    setError(false)
                        setUsetWalletAddress(accounts[0] ?? "")
                    
                    sessionStorage.setItem("account", wallet.accounts[0] as string)
                })
                .catch((err: any) => {
                    setError(true)
                    setErrorMessage(err.message)
                })
            setIsConnecting(false)
        }
        if(wordCoinAddress == null){
            getProvider()
            handleConnect()
        }else{
            sendTransaction()
        }
        
            return () => {
                window.ethereum?.removeListener('accountsChanged', refreshAccounts)
                window.ethereum?.removeListener("chainChanged", refreshChain)
            }
    },[wordCoinAddress])
    const abi = `
    [
        {
            "inputs": [
                {
                    "internalType": "contract IWorldID",
                    "name": "_worldId",
                    "type": "address"
                },
                {
                    "internalType": "string",
                    "name": "_appId",
                    "type": "string"
                },
                {
                    "internalType": "string",
                    "name": "_action",
                    "type": "string"
                }
            ],
            "stateMutability": "nonpayable",
            "type": "constructor"
        },
        {
            "inputs": [],
            "name": "InvalidNullifier",
            "type": "error"
        },
        {
            "anonymous": false,
            "inputs": [
                {
                    "indexed": false,
                    "internalType": "string",
                    "name": "message",
                    "type": "string"
                }
            ],
            "name": "MessageSent",
            "type": "event"
        },
        {
            "stateMutability": "payable",
            "type": "fallback"
        },
        {
            "inputs": [
                {
                    "internalType": "bytes",
                    "name": "encodedData",
                    "type": "bytes"
                }
            ],
            "name": "decodeData",
            "outputs": [
                {
                    "internalType": "uint256[8]",
                    "name": "",
                    "type": "uint256[8]"
                }
            ],
            "stateMutability": "pure",
            "type": "function"
        },
        {
            "inputs": [],
            "name": "mailbox",
            "outputs": [
                {
                    "internalType": "contract IMailbox",
                    "name": "",
                    "type": "address"
                }
            ],
            "stateMutability": "view",
            "type": "function"
        },
        {
            "inputs": [
                {
                    "internalType": "uint32",
                    "name": "_destinationDomain",
                    "type": "uint32"
                },
                {
                    "internalType": "string",
                    "name": "_message",
                    "type": "string"
                },
                {
                    "internalType": "address",
                    "name": "_recipientAddress",
                    "type": "address"
                }
            ],
            "name": "sendMessage",
            "outputs": [],
            "stateMutability": "payable",
            "type": "function"
        },
        {
            "inputs": [
                {
                    "internalType": "address",
                    "name": "signal",
                    "type": "address"
                },
                {
                    "internalType": "uint256",
                    "name": "root",
                    "type": "uint256"
                },
                {
                    "internalType": "uint256",
                    "name": "nullifierHash",
                    "type": "uint256"
                },
                {
                    "internalType": "bytes",
                    "name": "proof",
                    "type": "bytes"
                }
            ],
            "name": "verifyAndExecute",
            "outputs": [],
            "stateMutability": "nonpayable",
            "type": "function"
        },
        {
            "stateMutability": "payable",
            "type": "receive"
        }
    ]
    `
    const sendTransaction = async () => {
        try {
          // Request account access if needed
          await (window.ethereum as any).request({ method: 'eth_requestAccounts' });
    
          // Create a Web3 provider from MetaMask's provider
          const provider = new ethers.providers.Web3Provider(window.ethereum as any);
    
          // Get the signer from the provider
          const signer = provider.getSigner();
    
          // Create a new instance of the contract
          const contract = new ethers.Contract("0x5093a7Cab14f8d7Ea6F34AbfC8ee4b6535de53F5", abi, signer);
    
          // Send a transaction to the contract
          const tx = await contract["verifyAndExecute"](userWalletAddress,wordCoinAddress?.merkle_root,wordCoinAddress?.nullifier_hash,wordCoinAddress?.proof);
    
          // Wait for the transaction to be mined
          const receipt = await tx.wait();
    
          console.log('Transaction receipt:', receipt);
        } catch (error) {
          console.error('Transaction error:', error);
        }
      };
    const router = useRouter()
    
    const onSuccess = (data: ISuccessResult) => {
        console.log(data);
        setWordCoinAddress(data)
    }

    return (
        <div className="bg-white py-6 sm:py-8 lg:py-12">
            <div className="mx-auto max-w-lg">
                <IDKitWidget
                    app_id={process.env.NEXT_PUBLIC_WORLD_COIN_ID}// obtained from the Developer Portal
                    action="login_eth" // this is your action name from the Developer Portal
                    onSuccess={onSuccess} // callback when the modal is closed
                    signal={userWalletAddress} // prevents tampering with a message
                    credential_types={[CredentialType.Orb, CredentialType.Phone]} // optional, defaults to ['orb']
                >
                    {({ open }) =>
                        <div className="flex flex-col gap-4 p-4 md:p-8">
                            <button onClick={open} className="block rounded-lg bg-gradient-to-r from-green-400 via-green-500 to-green-600 px-8 py-3 text-center text-sm font-semibold text-white outline-none ring-gray-300 transition duration-100 hover:bg-gray-700 focus-visible:ring active:bg-gray-600 md:text-base">Verify with World ID</button>
                        </div>
                    }
                </IDKitWidget>

            </div>
        </div>
    )
}

export default Login;