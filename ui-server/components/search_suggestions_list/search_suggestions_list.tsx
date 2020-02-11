import { PodcastLink } from 'components/link'
import React from 'react'
import { SearchSuggestion } from 'types/models'
import { getImageUrl } from 'utils/dom'
import { uniqueId } from 'utils/utils'

export interface StateToProps {
  suggestions: SearchSuggestion[]
}

const SearchSuggestionsList: React.FC<StateToProps> = ({ suggestions }) => {
  if (suggestions.length === 0) {
    return <></>
  }

  return (
    <div
      style={{ width: '32rem' }}
      className="z-10 px-2 py-2 bg-white border border-blue-400 rounded-lg"
    >
      {suggestions.map((s) => {
        if (SearchSuggestion.isPodcast(s)) {
          return renderPodcast(s, uniqueId())
        }
        return renderText(s, uniqueId())
      })}
    </div>
  )
}

const renderPodcast = (p: SearchSuggestion, key: string): JSX.Element => (
  <PodcastLink key={key} podcastUrlParam={p.i}>
    <a className="search-suggestion block flex px-3 py-3 hover:bg-gray-200 rounded">
      <img src={getImageUrl(p.i)} className="w-12 h-12 mr-3 border rounded" />
      <div>
        <div
          className="text-base text-gray-900 line-clamp-1"
          dangerouslySetInnerHTML={{ __html: p.header }}
        />
        <div
          className="text-sm text-gray-900 line-clamp-1"
          dangerouslySetInnerHTML={{ __html: p.subHeader }}
        />
      </div>
    </a>
  </PodcastLink>
)

const renderText = (t: SearchSuggestion, key: string): JSX.Element => (
  <div key={key}>
    <div
      className="text-base text-gray-900 line-clamp-1"
      dangerouslySetInnerHTML={{ __html: t.header }}
    />
  </div>
)

export default SearchSuggestionsList
