import dynamic from "next/dynamic";
import Head from "next/head";
import Image from "next/image";
import React, { Suspense } from "react";
import { useState } from "react";
import ReactVisibilitySensor from "react-visibility-sensor";
import { Waypoint } from "react-waypoint";
import Carousel from "../components/Carousel";

const Listing = dynamic(() => import("../components/Listing"), {
  ssr: false,
});

export default function Home() {
  const [state, setState] = useState(false);
  return (
    <div>
      <Carousel />
      <Carousel />
      <Carousel />
      <Carousel />
      <Carousel />
      {Array.apply(null, Array(5)).map((_, idx) => (
        <Waypoint key={idx} onEnter={() => setState(true)}>
          <div>{state ? <Listing id={idx} /> : <>loading</>}</div>
        </Waypoint>
      ))}
    </div>
  );
}
