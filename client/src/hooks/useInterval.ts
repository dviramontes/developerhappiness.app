import { useEffect, useRef } from "react";

// attribution:
// https://gist.github.com/hyber1z0r/78e636ba95059cf2c86efdd1800897f3#file-useinterval-ts
export const useInterval = (callback: () => void, delay: number) => {
  const savedCallback = useRef(callback);

  useEffect(() => {
    savedCallback.current = callback;
  }, [callback]);

  useEffect(() => {
    function tick() {
      savedCallback.current();
    }
    if (delay !== null) {
      let id = setInterval(tick, delay);
      return () => clearInterval(id);
    }
  }, [delay]);
};
