"use client";

import { useEffect, useRef, useState } from "react";
import BasicCard from "../components/BasicInfoCard";

export default function NasaData() {
    //Data Storage for endpoint calls
    const [last50NasaData, setlast50NasaData] = useState([]);
    const [cardRows, setCardRows] = useState([])


    // Grab Nasa data (backend limits to 50 entries for demo)
    useEffect(()=>{
      const fetchAllNasaData = async () => {
        const allNasaDataResponse = await fetch("http://localhost:5050/api/v1/getAllNasaData");
        const allNasaDataJson = await allNasaDataResponse.json()
        console.log(allNasaDataJson)
        setlast50NasaData(allNasaDataJson.data)
      }

      fetchAllNasaData()

    }, [])

    // When nasa data is returned from backend, grab specific fields for data card
    useEffect(() => {
      let dataCardsTemp : any = []
      last50NasaData.forEach((nasaDataPoint : any) => {
        // For each data point, get name type, mass, year
        const dataPoints = [
          "NameType: " + nasaDataPoint.nameType,
          "Mass: " + nasaDataPoint.mass,
          "Year: " + nasaDataPoint.year
        ]

        // Create Data Card that contains Name, name type, mass, year
        // Also includes button to lead user to more details about a specific id
        dataCardsTemp.push(<BasicCard title={nasaDataPoint.name} data={dataPoints} buttonRoute={"/nasaData/"+nasaDataPoint.id}/>)
      })
      setCardRows(dataCardsTemp)
    }, [last50NasaData])


    //Page contains:
    // Header
    // Subheader
    // Cards for each data point
    // Button on each card to go to more details about specific data point (dynamic routes)
    return (
      <div className="p-4 inline-block">
      <h1 className="text-3xl font-bold pb-2">
        Earth Meteroite Landings
      </h1>
      <hr/>
      <h2 className="text-2xl font-bold underline pt-2 pb-2">
        Endpoint Responses
      </h2>
      {cardRows}
    </div>
    )
  }