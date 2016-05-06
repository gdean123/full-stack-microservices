require_relative './spec_helper'

describe 'to do list' do
  it 'renders hello world' do
    visit('/')
    expect(page).to have_content 'Hello, world!'
  end
end
