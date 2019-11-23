import React from 'react'
import Feed from './feed'

const HistoryView: React.FC<{}> = () => {
  console.log('iam rendering')
  return (
    <div className="flex md:flex-row flex-col">
      <div className="lg:w-2/3 w-full">
        <Feed />
      </div>
    </div>
  )
}

export default HistoryView
