import React from 'react'
import Head from 'next/head'
import axios from 'axios'
import PropTypes from 'prop-types'
import Navbar from '../components/Navbar'
import Body from '../components/Body'
import GlobalStyles from '../components/Styles/Global.styles'
import ApiUrlContext from '../context/API_URL.js'

function App({ users, API_URL }) {
  return (
    <ApiUrlContext.Provider value={API_URL}>
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
    </ApiUrlContext.Provider>
  )
}

App.getInitialProps = async () => {
  try {
    const connectionString = process.env.API_URL
      ? process.env.API_URL
      : 'https://api-dot-resume-engine.appspot.com'
    const { data } = await axios.get(`${connectionString}/api/user`)
    return { users: data, API_URL: connectionString }
  } catch (err) {
    console.error(err)
    throw err
  }
}

App.propTypes = {
  users: PropTypes.arrayOf(PropTypes.shape({})).isRequired,
  API_URL: PropTypes.string.isRequired,
}

export default App
