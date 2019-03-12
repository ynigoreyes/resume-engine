import React, { useEffect, useState } from 'react'
import { Grid, Dimmer, Loader, Image, Segment } from 'semantic-ui-react'
import axios from 'axios'
import CommentsContainer from './Comments/CommentsContainer'
import UserChipList from './UserChip/UserChipList'

function Body({ users }) {
  let newUsers = users
  const fetchComments = (id) => async () => {
    try {
      const { data } = await axios.get('http://localhost:8080/api/user')
      console.log(data)
    } catch (err) {
      console.error(err)
      newUsers = [{
        ID: 1,
        first_name: 'John',
        last_name: 'Doe',
        tag_line: 'tag',
      }]
    }
  }

  return (
    <main>
      <Grid columns={2}>
        <Grid.Column width={5}>
          <UserChipList fetchComments={fetchComments} users={newUsers} />
        </Grid.Column>
        <Grid.Column width={7}>
          <CommentsContainer />
        </Grid.Column>
      </Grid>
    </main>
  )
}

export default Body
