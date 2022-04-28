import { useState, useEffect } from "react";
import "./App.css";

function App() {
  const [data, setData] = useState<any>();

  useEffect(() => {
    getData();
    async function getData() {
      const result = await fetch("http://localhost:9090/products");
      const json = await result.json();
      setData(json);
    }
  }, []);

  return (
    <div className="App">
      {data ? <pre>{JSON.stringify(data, null, 2)}</pre> : "No Data"}
    </div>
  );
}

export default App;
