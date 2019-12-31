import ButtonSubscribe from 'components/button_subscribe'
import format from 'date-fns/format'
import React from 'react'
import { Podcast } from 'types/app'
import { getImageUrl } from 'utils/dom'

export interface OwnProps {
  podcast: Podcast
}

const PodcastInfo: React.SFC<OwnProps> = ({ podcast }) => {
  return (
    <div className="flex">
      <img
        className="lg:h-36 h-24 lg:w-36 w-24 flex-none object-contain object-center rounded-lg border"
        src={getImageUrl(podcast.urlParam)}
      />
      <div className="flex flex-col flex-auto w-1/2 justify-between lg:px-5 px-3">
        <div className="w-full">
          <h2 className="md:text-xl text-lg text-gray-900 leading-tight line-clamp-2">
            {podcast.title}
          </h2>
          <h3 className="md:text-base text-sm text-gray-800 leading-loose truncate">
            {podcast.author}
          </h3>
          <h4 className="text-xs text-gray-700">
            {`Since ${format(
              new Date(`${podcast.earliestEpisodePubDate} +0000`),
              'MMM yyyy',
            )}`}
            <span className="mx-2 font-extrabold">&middot;</span>
            {`${podcast.totalEpisodes} episodes`}
          </h4>
        </div>
        <ButtonSubscribe className="w-24 py-1 text-xs" podcastId={podcast.id} />
      </div>
    </div>
  )
}

export default PodcastInfo
