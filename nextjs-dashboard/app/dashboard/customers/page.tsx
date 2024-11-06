'use client'

import { useState } from "react"

export default function Page() {
    const [state, setState] = useState(null)

    async function handleClick() {
      let data = await fetch('http://localhost:3000/api/admin')
      let res = await data.json()
      setState(res.message)
    }

    return (
      <>
        <button onClick={handleClick}>Click</button>
        <p>{state}</p>
        <a href="/image-20241029213335791.png" download="image.png">link</a>
      </>
    )
}