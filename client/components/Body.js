import React, { useEffect, useState } from 'react'
import { Grid, Dimmer, Loader, Image, Segment } from 'semantic-ui-react'
import CommentsContainer from './Comments/CommentsContainer'
import UserChipList from './UserChip/UserChipList'

function Body({ users }) {
  const fetchComments = (id) => async () => {
    console.log(id)
  }

  return (
    <main>
      <Grid columns={2}>
        <Grid.Column width={5}>
          <UserChipList fetchComments={fetchComments} users={users} />
        </Grid.Column>
        <Grid.Column width={7}>
          <CommentsContainer />
        </Grid.Column>
      </Grid>
    </main>
  )
}

export default Body
