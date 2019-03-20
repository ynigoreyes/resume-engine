import React, { useState, useEffect, useContext } from 'react'
import { Form, TextArea, Button } from 'semantic-ui-react'
import { SCREENS } from '../Body'
import CommentsHeader from './CommentsHeader'
import Environment from '../../context/Environment'

const style = {
  form: {
    height: '100%',
    background: 'white',
    margin: '25px 0px',
    borderRadius: '50px 0px 0px 50px',
    overflow: 'hidden',
  },
  bar: {
    height: '10px',
    background: '#595a4a',
  },
  footer: {
    position: 'absolute',
    bottom: 0,
    height: '100px',
    padding: '30px 60px',
    background: '#595a4a',
    width: '100%',
  },
  textArea: {
    fontSize: '18px',
  },
}

function parseComments(comments) {
  let content = comments
  if (comments.length) {
    content = ''
    for (const comment of comments) {
      content += `"${comment.comment_body}"\n~ ${
        comment.commenter_first_name
      }, ${comment.commenter_last_name}\n\n`
    }
  } else {
    content = 'No comments have been given\nSorry!'
  }
  return content
}

function CommentsContainer({
  currentCommenter,
  currentResume,
  comments,
  screen,
  createComment,
}) {
  const env = useContext(Environment)
  const [commentArea, setCommentArea] = useState('')
  useEffect(() => {
    if (screen === SCREENS.COMMENTER) {
      setCommentArea('')
    } else {
      setCommentArea(parseComments(comments))
    }
  }, [screen, comments])

  return (
    <Form style={style.form}>
      <div style={style.bar} />
      <CommentsHeader
        imgUrl={
          screen === SCREENS.COMMENTER
            ? `${env.STORAGE_URL}/GCP.gif`
            : `${env.STORAGE_URL}/avatar_${currentResume.ID - 1}.gif`
        }
        name={
          screen === SCREENS.COMMENTER
            ? currentCommenter.first_name
            : currentResume.first_name
        }
        tag={
          screen === SCREENS.COMMENTER
            ? 'What do you think about this resume?'
            : 'Here is what people have said about you'
        }
      />
      {screen === SCREENS.COMMENTER ? (
        <TextArea
          value={commentArea}
          placeholder='Please be nice to your friends'
          rows={13}
          style={style.textArea}
          onChange={(e) => setCommentArea(e.target.value)}
        />
      ) : (
        <TextArea
          value={commentArea}
          rows={13}
          style={style.textArea}
          readOnly
        />
      )}
      <div style={style.footer}>
        {screen === SCREENS.COMMENTER ? (
          <>
            <Button
              style={{ background: '#F7F4F3' }}
              onClick={createComment(commentArea)}
              size='large'
            >
              Submit
            </Button>
            <Button
              basic
              size='large'
              inverted
              onClick={() => setCommentArea('')}
            >
              Clear
            </Button>
          </>
        ) : null}
      </div>
    </Form>
  )
}

export default CommentsContainer
