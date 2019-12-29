import React from 'react'
import Feed from './feed'

const PlaylistView: React.FC<{}> = () => {
  return (
    <div>
      <div className="lg:w-2/3 w-full lg:mr-5">
        <Feed />
      </div>
      <div className="lg:w-1/3 w-full lg:ml-3">
        
      </div>
    </div>
  )
}

export default PlaylistView
