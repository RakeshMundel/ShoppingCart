import React from 'react'
import Hero from '../Components/Hero/Hero';
import Popular from '../Components/Popular/Popular';

const Shop = () => {
  console.log("Shop rendering");

  return (
    <div>
      <Hero/>
      <hr style={{ background: "#e0e0e0", height: "2px", border: "none" }} />
      <Popular/>
    </div>
  )
}

export default Shop