import React from 'react'
import Head from 'next/head'
import axios from 'axios'
import Navbar from '../components/Navbar'
import Body from '../components/Body'
import GlobalStyles from '../components/Styles/Global.styles'

const initUsers = [{
  ID: 1,
  first_name: 'John',
  last_name: 'Doe',
  tag_line: 'tag',
}]

function App({ users }) {
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
      <Body users={users} />
    </main>
  )
}

App.getInitialProps = async () => {
  try {
    const { data } = await axios.get('https://api-dot-resume-engine.appspot.com/api/user')
    return { users: data }
  } catch (err) {
    console.error(err)
    return { users: initUsers }
  }
}
export default App
