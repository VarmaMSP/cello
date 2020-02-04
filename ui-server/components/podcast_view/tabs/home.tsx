import React from 'react'
import { Podcast } from 'types/models'
import CategoryList from '../components/category_list'
import EpisodeList from '../components/episode_list'

export interface OwnProps {
  podcast: Podcast
}

const PodcastAbout: React.FC<OwnProps> = ({ podcast }) => {
  return (
    <div>
      <h2 className="font-medium tracking-wider mb-2">{'Description'}</h2>
      <div
        className="text-gray-900 text-sm tracking-wide leading-relaxed"
        style={{ hyphens: 'auto' }}
      >
        <div>{podcast.description}</div>
        <div className="my-4">
          <CategoryList
            categoryIds={podcast.categories.map((x) => x.categoryId)}
          />
        </div>
        {!!podcast.copyright && (
          <div className="text-gray-700 line-clamp-1">{`${podcast.copyright}`}</div>
        )}
      </div>

      <hr className="my-6" />

      <h2 className="font-medium tracking-wider mb-5">{'Episodes'}</h2>
      <EpisodeList podcastId={podcast.id} />
    </div>
  )
}

export default PodcastAbout
