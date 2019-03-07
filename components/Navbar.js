import React from 'react'

const NavBarStyles = () => (
  <style jsx>
    {`
      div {
        background-color: black;
      }
    `}
  </style>
)

function Navbar() {
  return (
    <div>
      <NavBarStyles />
      Google Critiques
    </div>
  )
}

export default Navbar
