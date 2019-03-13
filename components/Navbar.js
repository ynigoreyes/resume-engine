import React from 'react'

const style = {
  height: '60px',
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'center',
}

function Navbar() {
  return (
    <div style={style} className='green'>
      <h3>Google Critiques</h3>
    </div>
  )
}

export default Navbar
