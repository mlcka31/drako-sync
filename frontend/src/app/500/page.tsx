"use client"; // Error boundaries must be Client Components

import { useEffect } from "react";

export default function Error() {
  return (
    <div>
      <h2>Something went wrong!</h2>
      <button>Try again</button>
    </div>
  );
}
