describe 'Commands' do
  describe 'down', command: true do
    it 'should clear out the IP alias' do
      loon %(up)
      ip = project_ip
      loon %(down)

      expect(`ifconfig`).not_to include(ip)
    end
  end
end

