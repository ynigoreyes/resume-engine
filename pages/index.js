import React from 'react'
import Head from 'next/head'
import Navbar from '../components/Navbar'
import Body from '../components/Body'
import GlobalStyles from '../components/Styles/Global.styles'

function App() {
  return (
    <main>
      <Head>
        <title>Google Cloud</title>
        <link
          rel='stylesheet'
          href='//cdn.jsdelivr.net/npm/semantic-ui@2.4.2/dist/semantic.min.css'
        />
      </Head>
      <GlobalStyles />
      <Navbar />
      <Body />
    </main>
  )
}
export default App
