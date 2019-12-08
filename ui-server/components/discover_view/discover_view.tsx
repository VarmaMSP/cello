import React from 'react'
import PodcastList from './podcast_list'

interface OwnProps {
  listId: string
}

const DiscoverView: React.FC<OwnProps> = ({ listId }) => {
  return (
    <div className="flex md:flex-row flex-col">
      <div className="md:w-2/3 w-fullpr-2">
      <PodcastList listId={listId} />
      </div>
    </div>
  )
}

export default DiscoverView
