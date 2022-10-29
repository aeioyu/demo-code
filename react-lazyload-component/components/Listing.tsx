import React from "react";

const Listing = ({ id }: { id: number }): JSX.Element => {
  console.log("rendering", id);
  return <div style={{ background: "#e0e1e2", height: 300, marginBottom: 10 }}>Listing {id}</div>;
};

export default Listing;
