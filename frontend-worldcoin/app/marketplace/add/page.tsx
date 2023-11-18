"use client"
import React, { useEffect, useState } from 'react';
import { ethers } from "ethers";
import { useSDK } from '@metamask/sdk-react';
import { Web3Storage } from 'web3.storage';
import detectEthereumProvider from '@metamask/detect-provider';


const AddhtmlForm = () => {
    // State variables for each input
    const [name, setName] = useState('');
    const [cid, setCid] = useState('');
    const [shortDescription, setShortDescription] = useState('');
    const [data, setData] = useState('');
    const [file, setFile] = useState(null);
    const [account, setAccount] = useState<string>();
    const { sdk, connected, connecting, provider, chainId } = useSDK();
  
    // Handlers for each input
    const handleNameChange = (e: any) => setName(e.target.value);
    const handleShortDescriptionChange = (e: any) => setShortDescription(e.target.value);
    const handleDataChange = (e: any) => setData(e.target.value);
    const handleFileChange = (e: any) => setFile(e.target.files[0]);

    const handleConnect = async () => {
        await window.ethereum.request({
            method: "eth_requestAccounts",
        })
            .then((_) => {
                
            })
            .catch((err: any) => {
                console.log(err.message)
            })
    }
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
            await handleConnect()
            // Create a Web3 provider from MetaMask's provider
            const provider = await detectEthereumProvider({ silent: true })
            const web3Provider = new ethers.BrowserProvider(window.ethereum);
            // Get the signer from the provider
            const signer = await web3Provider.getSigner();

            // Create a new instance of the contract
            const contract = new ethers.Contract("0x5E622554E119C8Dc90b4BE2ae2fF9D4fA034645C", abi, signer);
            const feeAmount = ethers.parseEther("0.01"); // Replace with the actual fee amount in Ether

            // Send a transaction to the contract
            const tx = await contract["sendMessage"](534351, data, '0x5093a7Cab14f8d7Ea6F34AbfC8ee4b6535de53F5',{ value: feeAmount });

            // Wait for the transaction to be mined
            const receipt = await tx.wait();

            console.log('Transaction receipt:', receipt);
        } catch (error) {
            console.error('Transaction error:', error);
        }
    };
    const buttonOnClick = async(event: React.MouseEvent<HTMLButtonElement>) => {
        const client = new Web3Storage({ token: process.env.NEXT_PUBLIC_STORAGE_TOKEN });
        const _name = new File([name], "name.txt", { type: "text/plain" });
        const _shortDescription = new File([shortDescription], "shortDescription.txt", { type: "text/plain" });
        const _data = new File([data], "data.txt", { type: "text/plain" });
    
        const rootCid = await client.put([_name, _shortDescription, _data]);
        setCid(rootCid)
        //sendTransaction()
    }

    return (
        <div className="bg-white py-6 sm:py-8 lg:py-12">
            <div className="mx-auto max-w-screen-2xl px-4 md:px-8">
                <div className="mb-10 md:mb-16">
                    <h2 className="mb-4 text-center text-2xl font-bold text-gray-800 md:mb-6 lg:text-3xl">Customize Your Own Protocol GPT</h2>

                </div>
                <div className="mx-auto grid max-w-screen-md gap-4 sm:grid-cols-2">

                    <div>
                        <label htmlFor="first-name" className="mb-2 inline-block text-sm text-gray-800 sm:text-base">Name</label>
                        <input
                            name="first-name"
                            className="w-full rounded border bg-gray-50 px-3 py-2 text-gray-800 outline-none ring-indigo-300 transition duration-100 focus:ring"
                            value={name}
                            onChange={handleNameChange}
                        />
                    </div>

                    <div>
                        <label htmlFor="last-name" className="mb-2 inline-block text-sm text-gray-800 sm:text-base">Short Description</label>
                        <input
                            name="last-name"
                            className="w-full rounded border bg-gray-50 px-3 py-2 text-gray-800 outline-none ring-indigo-300 transition duration-100 focus:ring"
                            value={shortDescription}
                            onChange={handleShortDescriptionChange}
                        />
                    </div>

                    <div className="sm:col-span-2">
                        <label htmlFor="message" className="mb-2 inline-block text-sm text-gray-800 sm:text-base">Data*</label>
                        <textarea
                            name="message"
                            className="h-64 w-full rounded border bg-gray-50 px-3 py-2 text-gray-800 outline-none ring-indigo-300 transition duration-100 focus:ring"
                            value={data}
                            onChange={handleDataChange}
                        ></textarea>
                    </div>

                    <div className="sm:col-span-2">
                        <label className="block mb-2 text-sm font-medium text-gray-900 dark:text-black" htmlFor="file_input">Upload Image</label>
                        <input
                            className="block w-full p-2 text-sm text-gray-900 border rounded-lg cursor-pointer bg-gray-50 focus:outline-none dark:placeholder-gray-400"
                            id="file_input"
                            type="file"
                            onChange={handleFileChange}
                        />
                    </div>
                    { cid &&
                        <div className="sm:col-span-2">
                            <label className="block mb-2 text-sm font-medium text-gray-900 dark:text-black" htmlFor="file_input">File is uploaded to <br/> {cid}</label>
                        </div>
                    }
                    <div className="flex items-center w-full justify-between sm:col-span-2">
                        <button onClick={buttonOnClick} className="inline-block w-full rounded-lg bg-indigo-500 px-8 py-3 text-center text-sm font-semibold text-white outline-none ring-indigo-300 transition duration-100 hover:bg-indigo-600 focus-visible:ring active:bg-indigo-700 md:text-base">Send</button>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default AddhtmlForm;
