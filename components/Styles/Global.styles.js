import React from 'react'

function GlobalStyles() {
  return (
    <style jsx global>
      {`
        body {
          background-color: #363635;
          margin: 0;
          padding: 0;
        }
        .green {
          background-color: #595a4a;
          color: #f7f4f3;
        }
        .dark-gray {
          background-color: #363635;
          color: #363635;
        }
        .gray {
          background-color: #564d4a;
          color: #595a4a;
        }
        .light-gray {
          background-color: #bebebe;
          color: #bebebe;
        }
        .white {
          background-color: #f7f4f3;
          color: #f7f4f3;
        }
        h1 {
          font-family: Rockwell;
          font-size: 32px;
          font-style: normal;
          font-variant: normal;
          font-weight: 700;
          line-height: 32px;
        }
        h3 {
          font-family: Rockwell;
          font-size: 25px;
          font-style: normal;
          font-variant: normal;
          font-weight: 700;
          line-height: 32px;
        }
        p {
          font-family: Verdana;
          font-size: 20px;
          font-style: normal;
          font-variant: normal;
          font-weight: 400;
          line-height: 32px;
        }
      `}
    </style>
  )
}

export default GlobalStyles
