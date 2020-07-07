describe 'Dependencies' do
  describe 'Node.JS', dependency: true, node: true do
    versions, default = versions_for :node

    versions.each do |version|
      it "installs v#{version} correctly" do
        test_dep 'node', version: version, match: version
      end
    end

    it "installs v#{default} as the default" do
      test_dep 'node', match: default
    end
  end
end
