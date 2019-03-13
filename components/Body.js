import React, { useEffect, useState } from 'react'
import { Dimmer, Loader, Image, Segment } from 'semantic-ui-react'
import CommentsContainer from './Comments/CommentsContainer'
import UserChipList from './UserChip/UserChipList'

function Body() {
  const [loading, setLoading] = useState(true)
  const [users, setUsers] = useState([])

  useEffect(() => {
    console.log('fetch users')
    setUsers([])
    setLoading(false)
  }, [])

  const fetchComments = () => {}

  return (
    <main>
      {!loading ? (
        <>
          <UserChipList fetchComments={fetchComments} users={users} />
          <CommentsContainer />
        </>
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
