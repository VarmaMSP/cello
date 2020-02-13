import React from 'react'
import { getAssetUrl } from 'utils/dom'

const HomeView: React.FC<{}> = () => {
  return (
    <div>
      <div className="flex pt-6">
        <div className="flex-1"></div>
        <img src={getAssetUrl('listen-to-podcasts.svg')} className="w-1/2" />
      </div>
    </div>
  )
}

export default HomeView
