import { PodcastLink } from 'components/link'
import React from 'react'
import { Podcast, PodcastList } from 'types/app'
import { getImageUrl } from 'utils/dom'

export interface StateToProps {
  list: PodcastList
  podcasts: Podcast[]
}

export interface OwnProps {
  listId: string
}

const ListPodcasts: React.FC<StateToProps & OwnProps> = ({
  list,
  podcasts,
}) => {
  return (
    <div>
      <h2 className="text-xl text-gray-700">{list.title}</h2>
      <hr className="mt-1 mb-4 border-gray-400" />
      <ol>
        {podcasts.map((p, i) => (
          <PodcastLink key={p.id} podcastId={p.id}>
            <a>
              <li className="block flex items-center py-2 my-1 hover:bg-gray-200 rounded-lg">
                <div className="ml-2 mr-4 text-gray-600">{`${i + 1}.`}</div>
                <img
                  className="w-16 h-16 mr-4 flex-none object-contain rounded border cursor-default"
                  src={getImageUrl(p.urlParam)}
                />
                <div className="h-16">
                  <div className="text-gray-900 tracking-wide leading-loose line-clamp-1">
                    {p.title}
                  </div>
                  <div className="text-sm text-gray-800 traking-wide leading-relaxed line-clamp-1">
                    {p.author}
                  </div>
                </div>
              </li>
            </a>
          </PodcastLink>
        ))}
      </ol>
    </div>
  )
}

export default ListPodcasts
