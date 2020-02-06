import { PodcastLink } from 'components/link'
import React from 'react'
import { PodcastSearchResult } from 'types/models'
import { getImageUrl } from 'utils/dom'

export interface StateToProps {
  podcasts: PodcastSearchResult[]
}

const SearchSuggestions: React.FC<StateToProps> = ({ podcasts }) => {
  if (podcasts.length === 0) {
    return <></>
  }

  return (
    <div
      style={{ width: '32rem' }}
      className="z-10 px-2 py-2 bg-white border rounded-lg shadow"
    >
      {podcasts.map((x) => (
        <PodcastLink key={x.urlParam} podcastUrlParam={x.urlParam}>
          <a className="search-suggestion block flex px-3 py-3 hover:bg-gray-200 rounded">
            <img
              src={getImageUrl(x.urlParam)}
              className="w-12 h-12 mr-3 border rounded"
            />
            <div>
              <div
                className="text-base text-gray-800 line-clamp-1"
                dangerouslySetInnerHTML={{ __html: x.title }}
              />
              <div
                className="text-sm text-gray-700 line-clamp-1"
                dangerouslySetInnerHTML={{ __html: x.author }}
              />
            </div>
          </a>
        </PodcastLink>
      ))}
    </div>
  )
}

export default SearchSuggestions
