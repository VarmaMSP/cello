import { Podcast } from 'models'
import React from 'react'
import CategoryList from '../components/category_list'
import EpisodeList from '../components/episode_list'

export interface OwnProps {
  podcast: Podcast
}

const PodcastAbout: React.FC<OwnProps> = ({ podcast }) => {
  return (
    <div>
      <div
        className="mt-6 text-gray-800 text-sm font-medium leading-relaxed"
        style={{ hyphens: 'auto' }}
      >
        <div>{podcast.description}</div>
        <div className="mt-4">
          <CategoryList
            categoryIds={podcast.categories.map((x) => x.categoryId)}
          />
        </div>
      </div>

      <hr className="my-6" />

      <h2 className="font-medium tracking-wide mb-5">{'Episodes'}</h2>
      <EpisodeList podcastId={podcast.id} />
    </div>
  )
}

export default PodcastAbout
