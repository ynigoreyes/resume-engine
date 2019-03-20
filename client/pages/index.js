import React from 'react'
import Head from 'next/head'
import axios from 'axios'
import PropTypes from 'prop-types'
import Navbar from '../components/Navbar'
import Body from '../components/Body'
import GlobalStyles from '../components/Styles/Global.styles'
import Environment from '../context/Environment'

function App({ users, API_URL, API_SERVICE_NAME, STORAGE_URL }) {
  const env = {
    API_URL,
    API_SERVICE_NAME,
    STORAGE_URL,
  }

  return (
    <Environment.Provider value={env}>
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
        <Body users={users} />
      </main>
    </Environment.Provider>
  )
}

App.getInitialProps = async () => {
  try {
    const connectionString = process.env.NODE_ENV !== 'production'
      ? process.env.API_URL
      : `https://${process.env.API_SERVICE_NAME}-dot-resume-engine.appspot.com`
    const { data } = await axios.get(`${connectionString}/api/user`)
    return {
      users: data,
      API_URL: connectionString,
      API_SERVICE_NAME: process.env.API_SERVICE_NAME,
      STORAGE_URL: process.env.STORAGE_URL,
    }
  } catch (err) {
    console.error(err)
    throw err
  }
}

App.propTypes = {
  users: PropTypes.arrayOf(PropTypes.shape({})).isRequired,
  API_URL: PropTypes.string.isRequired,
  API_SERVICE_NAME: PropTypes.string.isRequired,
  STORAGE_URL: PropTypes.string.isRequired,
}

export default App
