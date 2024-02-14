"use client";

import { useEffect, useState } from "react"
import Button from '@mui/material/Button';
import { useRouter } from 'next/navigation';
import BasicCard from "@/app/components/BasicInfoCard";

export default function NasaDataForId(props: any) {
  console.log(props.params)

  //Data Storage for endpoint calls
  const [nasaDataPoint, setNasaDataPoint] = useState([]);
  const [cardRows, setCardRows] = useState([])
  const { push } = useRouter();


  // Grab Nasa data for specific url id
  useEffect(()=>{
    const fetchNasaDataPoint = async () => {
      // Use props.params.id (id is created from folder structure [id])
      const nasaDataResponse = await fetch("http://localhost:5050/api/v1/getNasaData/" + props.params.id );
      const nasaDataJson = await nasaDataResponse.json()
      console.log(nasaDataJson)
      setNasaDataPoint(nasaDataJson.data)
    }

    fetchNasaDataPoint()

  }, [])

  useEffect(() => {
    let dataCardsTemp : any = [] 
    nasaDataPoint.forEach((nasaDataForId : any) => {
      // Pull data needed for each data card and create card
      const dataPoints = [
        "NameType: " + nasaDataForId.nameType,
        "Mass: " + nasaDataForId.mass,
        "Year: " + nasaDataForId.year,
        "Classification: " + nasaDataForId.recclass,
        "Fall: " + nasaDataForId.fall,
        "Latitude: " + nasaDataForId.reclat,
        "Longitude: " + nasaDataForId.reclong,
      ]
      dataCardsTemp.push(<BasicCard title={nasaDataForId.name} data={dataPoints}/>)
    })
    setCardRows(dataCardsTemp)
  }, [nasaDataPoint])

  // Dynamic page title for nasa data name (needs data to be loaded to use name as title)
  let pageTitle = "NO DATA FOR ID"
  if (nasaDataPoint.length > 0){
    const dataForId: any = nasaDataPoint[0]
    pageTitle = dataForId.name
  }
 
    //Page contains:
    // Header
    // Subheader
    // Cards for each data point (in this case should be one card because we are viewing data about one id, we could of not used a list but this makes it more reusable)
    // Button at bottom to return to page that shows all data points
  return (
    <div className="p-4 inline-block">
      <h1 className="text-3xl font-bold pb-2">
        Earth Meteroite Landings
      </h1>
      <hr/>
      <h2 className="text-2xl font-bold underline pt-2 pb-2">
        {pageTitle}
      </h2>
      {cardRows}
      <Button variant="outlined" className="mt-2 mb-2" onClick={() => {push("/nasaData")}}>Check out Nasa Data here!</Button>
    </div>
  )
}