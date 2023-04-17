import Head from 'next/head'
import Image from 'next/image'
import { Inter } from 'next/font/google'
import styles from '@/styles/Home.module.css'
import { Container, Row } from 'react-bootstrap'
import SideBar from './components/sidebox/SideBar'
import ListOval from './components/sidebox/namebox/ListOval'
import MainBox from './components/mainbox/MainBox'
const inter = Inter({ subsets: ['latin'] })
const testprop = [
  {
    "company_name": "カイシャ1",
    "company_id": "kaisya1",
    "user_id": "mono",
    "recruitment_status": "1"
  },
  {
    "company_name": "カイシャ2",
    "company_id": "kaisya2",
    "user_id": "mono",
    "recruitment_status": "2"
  },
  {
    "company_name": "カイシャ3",
    "company_id": "kaisya3",
    "user_id": "mono",
    "recruitment_status": "3"
  },
  {
    "company_name": "カイシャ4",
    "company_id": "kaisya4",
    "user_id": "mono",
    "recruitment_status": "4"
  }

]


export default function Home() {
  return (
    <>
      <Head>
        <title>就活メモ</title>
        <meta name="description" content="Generated by create next app" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <div>


      <SideBar datas={testprop}/> 
      <MainBox/>

      </div>
    </>
  )
}
