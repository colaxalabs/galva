import React, { useEffect, useState } from 'react';
import './App.css';

function App() {
  // do we have an ethereum browser
  const [isEthereumWindow, setIsEthereumWindow] = useState<boolean>(false);

  useEffect(() => {
    // @ts-ignore
    setIsEthereumWindow(window.ethereum !== undefined);
  }, [isEthereumWindow])

  return (
    <div>
      Galva
    </div>
  );
}

export default App;
