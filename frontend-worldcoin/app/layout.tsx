import type { Metadata } from 'next'
import { Inter } from 'next/font/google'
import './globals.css'
import CustomNavbar from '@/components/navbar'

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
    <html className=' h-full' lang="en">
      <body className={`bg-primary min-h-screen min-w-screen  bg-white ${inter.className}`}>
        <CustomNavbar />
        <main className='h-screen bg-white'>
{children}</main>
      </body>
    </html>
  )
}
