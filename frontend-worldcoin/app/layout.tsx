import type { Metadata } from 'next'
import { Inter } from 'next/font/google'
import './globals.css'
import CustomNavbar from '@/components/navbar'
import { MetaMaskProvider } from '@metamask/sdk-react';
import { MetamaskProvider } from '@/hook/useMetamask';

const inter = Inter({ subsets: ['latin'] })

export const metadata: Metadata = {
  title: 'Welcome to Chat',
  description: 'New generation chating',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html className='h-full' lang="en">
      <MetamaskProvider>

      <body className={`bg-primary min-h-screen min-w-screen bg-white ${inter.className}`}>
        <CustomNavbar />
        <main className='h-screen px-10 py-10 bg-white'>
{children}</main>
      </body>
      </MetamaskProvider>
    </html>
  )
}
