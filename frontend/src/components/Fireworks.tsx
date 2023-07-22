import React, { useEffect, useRef } from 'react';
import { Fireworks } from 'fireworks-js';

const FireworksComponent = () => {
  // create the ref object with initial value null
  const containerRef = useRef(null);

  useEffect(() => {
    // containerRef.current will contain the actual DOM node (div element) after the first render
    if (containerRef.current) {
      const fireworks = new Fireworks(containerRef.current);
      fireworks.start();

      return () => {
        fireworks.stop();
      }
    }
  }, []);

  // assign the ref object to the div element
  // React will set containerRef.current to the corresponding DOM node
  return <div id="fireworks" ref={containerRef} />;
}

export default FireworksComponent;
