describe 'Dependencies' do
  describe 'Crystal', dependency: true, crystal: true do
    def test_crystal_dep(version: nil, match:)
      dep = if version
              {'crystal' => version}
            else
              'crystal'
            end

      with_payload(deps: dep) do |project|
        loon %w(up), dir: project
        loon %w(exec crystal --version), dir: project

        assert_stderr_empty
        assert_stdout match
        assert_status 0
      end
    end

    %w(0.35.1).each do |version|
      it "installs version #{version} correctly" do
        test_crystal_dep version: version, match: version
      end
    end

    it 'installs version 0.35.1 as the default' do
      test_crystal_dep match: '0.35.1'
    end
  end
end
