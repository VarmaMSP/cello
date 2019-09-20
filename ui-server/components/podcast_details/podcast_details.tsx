import { imageUrl } from 'components/utils'
import React, { useState } from 'react'
import Shiitake from 'shiitake'
import { Podcast } from 'types/app'
import SubscribeButton from './components/subscribe_button'

export interface StateToProps {
  podcast: Podcast
}

export interface OwnProps {
  podcastId: string
}

interface Props extends StateToProps, OwnProps {}

const PodcastDetails: React.SFC<Props> = ({ podcast }) => {
  const [lineCountDesc, setLineCountDesc] = useState(2)

  return (
    <div className="flex mb-8">
      <img
        className="lg:h-56 lg:w-56 h-36 w-36 flex-none object-contain object-center rounded-lg border"
        src={imageUrl(podcast.id, 'md')}
      />
      <div className="flex flex-col justify-between lg:px-5 px-3">
        <div>
          <h2 className="md:text-2xl text-lg text-gray-900">{podcast.title}</h2>
          <h3 className="md:text-lg text-base text-gray-800 leading-relaxed">
            {podcast.author}
          </h3>
          <Shiitake
            lines={lineCountDesc}
            throttleRate={200}
            renderFullOnServer
            className="lg:block hidden mt-1 text-sm text-gray-800"
            overflowNode={
              <span
                onClick={() => setLineCountDesc(lineCountDesc == 4 ? 100 : 4)}
                className="text-blue-700 cursor-pointer"
              >
                {' ...read more'}
              </span>
            }
          >
            {podcast.description}
          </Shiitake>
        </div>
        <SubscribeButton
          className="md:w-32 w-24 md:h-10 h-8 md:text-base text-sm"
          podcastId={podcast.id}
        />
      </div>
    </div>
  )
}

export default PodcastDetails
