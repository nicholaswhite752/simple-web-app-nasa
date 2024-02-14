"use client";

import { useEffect, useState } from "react"
import BasicCard from "../components/BasicInfoCard"
import Button from '@mui/material/Button';
import { useRouter } from 'next/navigation';

export default function Home() {
  //Data Storage for endpoint calls
  const [welcomeMessage, setWelcomeMessage] = useState("Loading Nasa Data...");
  const { push } = useRouter();


  // Endpoint calls for basic endpoints
  useEffect(()=>{
    const fetchWelcomeData = async () => {
      const welcomeResponse = await fetch("http://localhost:5050/api/v1/loadNasaData");
      const welcomeResponseData = await welcomeResponse.json()
      setWelcomeMessage("Nasa Data Loaded!")
    }

    fetchWelcomeData()

  }, [])


  //Page contains
  // Header
  // Subheader
  // Card with basic welcome endpoint response
  // Button to go to page that displays NasaData from go backend
  return (
    <div className="p-4 inline-block">
      <h1 className="text-3xl font-bold pb-2">
        Web Api Testing Ground
      </h1>
      <hr/>
      <h2 className="text-2xl font-bold underline pt-2 pb-2">
        Endpoint Response
      </h2>
      <BasicCard title="Basic Welcome Endpoint" data={["Reponse: " + welcomeMessage]}/>
      <Button variant="outlined" className="mt-2 mb-2" onClick={() => {push("/nasaData")}}>Check out Nasa Data here!</Button>
    </div>

  )
}