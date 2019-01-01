# frozen_string_literal: true

require 'httparty'
require 'nokogiri'
require 'json'

def story_from_flex_item(flex)
  headline = flex.css('div.headline')
  headline_link = headline.css('a')
  return nil if headline_link.empty?

  url = headline_link.attribute('href').value
  text = headline_link.text
  text = nil if text.empty?
  blurb = flex.css('div.blurb').text
  blurb = nil if blurb.empty?
  story = { url: url, headline: text, blurb: blurb }
  story
end

def scrape_homepage(url = 'https://www.washingtonpost.com')
  html = HTTParty.get(url)

  page = Nokogiri::HTML(html)

  main_content = page.css('#main-content')

  stories = []

  main_content.css('div.flex-item').each do |flex|
    story = story_from_flex_item(flex)
    next if story.nil?

    stories << story
  end
  stories
end

puts JSON.dump(scrape_homepage)
