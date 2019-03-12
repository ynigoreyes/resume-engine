import React, { useEffect, useState } from 'react'
import { Grid, Dimmer, Loader, Image, Segment } from 'semantic-ui-react'
import axios from 'axios'
import CommentsContainer from './Comments/CommentsContainer'
import UserChipList from './UserChip/UserChipList'

function Body() {
  const [loading, setLoading] = useState(true)
  const [users, setUsers] = useState([])

  useEffect(() => {
    const fetchUsers = async () => {
      try {
        const { data } = await axios.get('http://localhost:8080/api/user')
        setUsers(data)
        setLoading(false)
      } catch (err) {
        console.error(err)
        setUsers([])
        setLoading(false)
      }
    }

    fetchUsers()
  }, [])

  const fetchComments = (id) => () => {
    console.log(id)
  }

  return (
    <main>
      {!loading ? (
        <Grid columns={2}>
          <Grid.Column width={5}>
            <UserChipList fetchComments={fetchComments} users={users} />
          </Grid.Column>
          <Grid.Column width={7}>
            <CommentsContainer />
          </Grid.Column>
        </Grid>
      ) : (
        <Segment>
          <Dimmer active inverted>
            <Loader inverted>Loading</Loader>
          </Dimmer>

          <Image src='https://react.semantic-ui.com/images/wireframe/short-paragraph.png' />
        </Segment>
      )}
    </main>
  )
}

export default Body
