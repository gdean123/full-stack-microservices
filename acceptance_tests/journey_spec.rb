require_relative './spec_helper'

describe 'to do list' do
  it 'renders the list of tasks' do
    visit('/')

    expect(page).to have_content 'Destroy the ring in Mordor'
    expect(page).to have_content 'Bring balance to the force'
    expect(page).to have_content 'Save Princess Peach... again'
    expect(page).to have_content 'Destroy 7 horcruxes'
    expect(page).to have_content 'Live long and prosper'
  end
end
