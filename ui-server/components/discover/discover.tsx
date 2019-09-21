import React from 'react'
import CurationView from './curation'
import Trending from './trending'

export interface StateToProps {
  curationIds: string[]
}

const Discover: React.SFC<StateToProps> = (props) => {
  const { curationIds } = props

  return (
    <div>
      <Trending />
      {curationIds.map((id) => (
        <CurationView curationId={id} key={id} />
      ))}
    </div>
  )
}

export default Discover
