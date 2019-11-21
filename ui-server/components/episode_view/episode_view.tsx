import Link from 'next/link'
import React, { useEffect, useRef } from 'react'
import { Episode, Podcast } from 'types/app'
import { getImageUrl } from 'utils/dom'

export interface StateToProps {
  episode: Episode
  podcast: Podcast
}

export interface OwnProps {
  episodeId: string
}

const EpisodeView: React.FC<StateToProps & OwnProps> = ({
  episode,
  podcast,
}) => {
  const ref = useRef(null) as React.RefObject<HTMLDivElement>

  useEffect(() => {
    if (ref.current) {
      const a = ref.current.getElementsByTagName('a')
      for (let i = 0; i < a.length; ++i) {
        a[i].setAttribute('target', '_blank')
      }

      const img = ref.current.getElementsByTagName('img')
      for (let i = 0; i < img.length; ++i) {
        img[i].remove()
      }
    }
  })

  return (
    <div className="flex md:flex-row flex-col">
      <div className="lg:w-2/3 w-full">
        <div className="flex">
          <img
            className="lg:h-36 h-24 lg:w-36 w-24 flex-none object-contain object-center rounded-lg border"
            src={getImageUrl(podcast.id, 'md')}
          />
          <div className="flex flex-col flex-auto w-1/2 justify-between lg:px-5 px-3">
            <div className="w-full">
              <div className="w-16 mb-2 text-center text-2xs leading-relaxed tracking-wider bg-gray-300 rounded-full">
                Episode
              </div>
              <h2 className="text-lg text-gray-900 leading-tight line-clamp-2">
                {episode.title}
              </h2>
              <Link
                href={{
                  pathname: '/podcasts',
                  query: { podcastId: podcast.id, activeTab: 'episodes' },
                }}
                as={`/podcasts/${podcast.id}/episodes`}
                key={podcast.id}
              >
                <a className="md:text-base text-sm text-gray-800 hover:text-gray-900 leading-loose truncate">
                  {podcast.title}
                </a>
              </Link>
            </div>
          </div>
        </div>
        <div className="mt-8">
          <div className="mb-2 tracking-wider">Description</div>
          <div
            ref={ref}
            className="external-html lg:pr-16 text-sm leading-relaxed tracking-wide text-gray-800"
            dangerouslySetInnerHTML={{ __html: episode.description }}
          />
        </div>
      </div>
    </div>
  )
}

export default EpisodeView
