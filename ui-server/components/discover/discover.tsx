import React from 'react'
import CurationView from './curation'

export interface StateToProps {
  curationIds: string[]
}

const Discover: React.SFC<StateToProps> = (props) => {
  const { curationIds } = props

  return (
    <div>
      {curationIds.map((id) => (
        <CurationView curationId={id} key={id} />
      ))}
    </div>
  )
}

export default Discover
