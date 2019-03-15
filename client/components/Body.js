import React from 'react'
import PropTypes from 'prop-types'
import { Grid } from 'semantic-ui-react'
import CommentsContainer from './Comments/CommentsContainer'
import UserChipList from './UserChip/UserChipList'

function Body({ users }) {
  const fetchComments = (id) => async () => {
    console.log(`Fetch comments with commenterId: ${id}`)
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

Body.propTypes = {
  // Users prefetched
  users: PropTypes.arrayOf(
    PropTypes.shape({
      first_name: PropTypes.string.isRequired,
      last_name: PropTypes.string.isRequired,
      tag_line: PropTypes.string.isRequired,
    }),
  ),
}

export default Body
